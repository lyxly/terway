//go:build privileged

package utils

import (
	"fmt"
	"net"
	"runtime"
	"testing"

	"github.com/AliyunContainerService/terway/pkg/tc"
	terwayTypes "github.com/AliyunContainerService/terway/types"
	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/containernetworking/plugins/pkg/testutils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
	"github.com/vishvananda/netlink"
)

func TestEnsureVlanUntagger(t *testing.T) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	var err error
	hostNS, err := testutils.NewNS()
	assert.NoError(t, err)

	containerNS, err := testutils.NewNS()
	assert.NoError(t, err)

	err = hostNS.Set()
	assert.NoError(t, err)

	defer func() {
		err := containerNS.Close()
		assert.NoError(t, err)

		err = testutils.UnmountNS(containerNS)
		assert.NoError(t, err)

		err = hostNS.Close()
		assert.NoError(t, err)

		err = testutils.UnmountNS(hostNS)
		assert.NoError(t, err)
	}()

	err = netlink.LinkAdd(&netlink.Dummy{
		LinkAttrs: netlink.LinkAttrs{Name: "eni"},
	})
	assert.NoError(t, err)
	eni, err := netlink.LinkByName("eni")
	assert.NoError(t, err)

	err = EnsureVlanUntagger(eni)
	if err != nil {
		t.Errorf("error ensure vlan untagger, %v", err)
		t.Fail()
	}
}

var _ = Describe("Test TC filter", func() {
	var hostNS ns.NetNS
	const nicName = "eni"
	BeforeEach(func() {
		var err error
		hostNS, err = testutils.NewNS()
		Expect(err).NotTo(HaveOccurred())

		err = hostNS.Do(func(netNS ns.NetNS) error {
			defer GinkgoRecover()
			fmt.Printf("LinkAdd ---------- \n")
			return netlink.LinkAdd(&netlink.GenericLink{
				LinkAttrs: netlink.LinkAttrs{
					Name:        nicName,
					TxQLen:      1000,
					NumTxQueues: 2,
					NumRxQueues: 2,
				},
				LinkType: "netdevsim",
			})
		})
		Expect(err).NotTo(HaveOccurred())
	})
	AfterEach(func() {
		_ = hostNS.Do(func(netNS ns.NetNS) error {
			defer GinkgoRecover()
			eni, err := netlink.LinkByName(nicName)
			Expect(err).NotTo(HaveOccurred())
			err = netlink.LinkDel(eni)
			Expect(err).NotTo(HaveOccurred())
			return nil
		})
		Expect(hostNS.Close()).To(Succeed())
		Expect(testutils.UnmountNS(hostNS)).To(Succeed())
	})

	It("add qdisc mq", func() {
		var err error
		err = hostNS.Do(func(netNS ns.NetNS) error {
			defer GinkgoRecover()
			eni, err := netlink.LinkByName(nicName)
			Expect(err).NotTo(HaveOccurred())

			err = EnsureMQQdisc(eni)
			Expect(err).NotTo(HaveOccurred())

			qds, err := netlink.QdiscList(eni)
			Expect(err).NotTo(HaveOccurred())

			found := false
			for _, qd := range qds {
				if qd.Type() != "mq" {
					continue
				}
				Expect(qd.Attrs().Parent).Should(Equal(uint32(netlink.HANDLE_ROOT)))
				Expect(qd.Attrs().Handle).Should(Equal(netlink.MakeHandle(1, 0)))
				found = true
				break
			}
			Expect(found).Should(BeTrue())
			return nil
		})
		Expect(err).NotTo(HaveOccurred())
	})
	It("set egress filter", func() {
		var err error
		err = hostNS.Do(func(netNS ns.NetNS) error {
			defer GinkgoRecover()

			eni, err := netlink.LinkByName(nicName)
			Expect(err).NotTo(HaveOccurred())

			err = SetEgressPriority(eni, netlink.MakeHandle(1, 1), &terwayTypes.IPNetSet{
				IPv4: &net.IPNet{
					IP:   net.ParseIP("192.168.1.1"),
					Mask: net.CIDRMask(32, 32),
				},
				IPv6: &net.IPNet{
					IP:   net.ParseIP("fd00::1"),
					Mask: net.CIDRMask(128, 128),
				},
			})
			Expect(err).NotTo(HaveOccurred())

			qds, err := netlink.QdiscList(eni)
			Expect(err).NotTo(HaveOccurred())

			for _, qd := range qds {
				_, ok := qd.(*netlink.Prio)
				if !ok {
					continue
				}

				u32, err := tc.FilterBySrcIP(eni, qd.Attrs().Handle, &net.IPNet{
					IP:   net.ParseIP("192.168.1.1"),
					Mask: net.CIDRMask(32, 32),
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(u32).NotTo(BeNil(), "tc filter with src ipv4 should be found. Qdisc %#v", qd)

				u32v6, err := tc.FilterBySrcIP(eni, qd.Attrs().Handle, &net.IPNet{
					IP:   net.ParseIP("fd00::1"),
					Mask: net.CIDRMask(128, 128),
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(u32v6).NotTo(BeNil(), "tc filter with src ipv6 should be found")
			}
			return nil
		})
		Expect(err).NotTo(HaveOccurred())
	})
})
