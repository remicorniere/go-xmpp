package xmpp // import "gosrc.io/xmpp"

import (
	"testing"
)

func TestValidJids(t *testing.T) {
	tests := []struct {
		jidstr   string
		expected Jid
	}{
		{jidstr: "test@domain.com", expected: Jid{"test", "domain.com", ""}},
		{jidstr: "test@domain.com/resource", expected: Jid{"test", "domain.com", "resource"}},
		// resource can contain '/' or '@'
		{jidstr: "test@domain.com/a/b", expected: Jid{"test", "domain.com", "a/b"}},
		{jidstr: "test@domain.com/a@b", expected: Jid{"test", "domain.com", "a@b"}},
		{jidstr: "domain.com", expected: Jid{"", "domain.com", ""}},
	}

	for _, tt := range tests {
		jid, err := NewJid(tt.jidstr)
		if err != nil {
			t.Errorf("could not parse correct jid: %s", tt.jidstr)
			continue
		}

		if jid == nil {
			t.Error("jid should not be nil")
		}

		if jid.Node != tt.expected.Node {
			t.Errorf("incorrect jid Node (%s): %s", tt.expected.Node, jid.Node)
		}

		if jid.Node != tt.expected.Node {
			t.Errorf("incorrect jid domain (%s): %s", tt.expected.Domain, jid.Domain)
		}

		if jid.Resource != tt.expected.Resource {
			t.Errorf("incorrect jid resource (%s): %s", tt.expected.Resource, jid.Resource)
		}
	}
}

func TestIncorrectJids(t *testing.T) {
	badJids := []string{
		"",
		"user@",
		"@domain.com",
		"user:name@domain.com",
		"user<name@domain.com",
		"test@domain.com@otherdomain.com",
		"test@domain com/resource",
	}

	for _, sjid := range badJids {
		if _, err := NewJid(sjid); err == nil {
			t.Error("parsing incorrect jid should return error: " + sjid)
		}
	}
}
