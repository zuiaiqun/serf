package command

import (
	"github.com/mitchellh/cli"
	"strings"
	"testing"
)

func TestMembersCommand_implements(t *testing.T) {
	var _ cli.Command = &MembersCommand{}
}

func TestMembersCommandRun(t *testing.T) {
	a1 := testAgent(t)
	defer a1.Shutdown()
	rpcAddr, ipc := testIPC(t, a1)
	defer ipc.Shutdown()

	ui := new(cli.MockUi)
	c := &MembersCommand{Ui: ui}
	args := []string{"-rpc-addr=" + rpcAddr}

	code := c.Run(args)
	if code != 0 {
		t.Fatalf("bad: %d. %#v", code, ui.ErrorWriter.String())
	}

	if !strings.Contains(ui.OutputWriter.String(), a1.SerfConfig().NodeName) {
		t.Fatalf("bad: %#v", ui.OutputWriter.String())
	}
}

func TestMembersCommandRun_statusFilter(t *testing.T) {
	a1 := testAgent(t)
	defer a1.Shutdown()

	ui := new(cli.MockUi)
	c := &MembersCommand{Ui: ui}
	args := []string{
		"-rpc-addr=" + a1.RPCAddr,
		"-status=a.*e",
	}

	code := c.Run(args)
	if code != 0 {
		t.Fatalf("bad: %d. %#v", code, ui.ErrorWriter.String())
	}

	if !strings.Contains(ui.OutputWriter.String(), a1.SerfConfig.NodeName) {
		t.Fatalf("bad: %#v", ui.OutputWriter.String())
	}
}

func TestMembersCommandRun_statusFilter_failed(t *testing.T) {
	a1 := testAgent(t)
	defer a1.Shutdown()

	ui := new(cli.MockUi)
	c := &MembersCommand{Ui: ui}
	args := []string{
		"-rpc-addr=" + a1.RPCAddr,
		"-status=(fail|left)",
	}

	code := c.Run(args)
	if code != 0 {
		t.Fatalf("bad: %d. %#v", code, ui.ErrorWriter.String())
	}

	if strings.Contains(ui.OutputWriter.String(), a1.SerfConfig.NodeName) {
		t.Fatalf("bad: %#v", ui.OutputWriter.String())
	}
}

func TestMembersCommandRun_roleFilter(t *testing.T) {
	a1 := testAgent(t)
	defer a1.Shutdown()

	ui := new(cli.MockUi)
	c := &MembersCommand{Ui: ui}
	args := []string{
		"-rpc-addr=" + a1.RPCAddr,
		"-role=test",
	}

	code := c.Run(args)
	if code != 0 {
		t.Fatalf("bad: %d. %#v", code, ui.ErrorWriter.String())
	}

	if !strings.Contains(ui.OutputWriter.String(), a1.SerfConfig.NodeName) {
		t.Fatalf("bad: %#v", ui.OutputWriter.String())
	}
}

func TestMembersCommandRun_roleFilter_failed(t *testing.T) {
	a1 := testAgent(t)
	defer a1.Shutdown()

	ui := new(cli.MockUi)
	c := &MembersCommand{Ui: ui}
	args := []string{
		"-rpc-addr=" + a1.RPCAddr,
		"-role=primary",
	}

	code := c.Run(args)
	if code != 0 {
		t.Fatalf("bad: %d. %#v", code, ui.ErrorWriter.String())
	}

	if strings.Contains(ui.OutputWriter.String(), a1.SerfConfig.NodeName) {
		t.Fatalf("bad: %#v", ui.OutputWriter.String())
	}
}
