package mock

import (
	"github.com/google/nftables"
	"github.com/sbezverk/nftableslib"
)

// Mock defines type and methods to simulate operations with tables
type Mock struct {
	ti nftableslib.TablesInterface
}

// Flush returns
func (m *Mock) Flush() error {
	_, err := m.ti.Tables().Dump()
	if err != nil {
		return err
	}
	return nil
}

// FlushRuleset not use
func (m *Mock) FlushRuleset() {

}

// AddRule not use
func (m *Mock) AddRule(r *nftables.Rule) *nftables.Rule {
	return r
}

// DelRule not used
func (m *Mock) DelRule(*nftables.Rule) error {
	return nil
}

// DelTable not used
func (m *Mock) DelTable(t *nftables.Table) {
}

// AddTable not used
func (m *Mock) AddTable(t *nftables.Table) *nftables.Table {
	return t
}

// AddChain not used
func (m *Mock) AddChain(c *nftables.Chain) *nftables.Chain {
	return c
}

// DelChain not used
func (m *Mock) DelChain(c *nftables.Chain) {
}

// AddSet not used
func (m *Mock) AddSet(s *nftables.Set, se []nftables.SetElement) error {
	return nil
}

// GetRule not implemented yet
func (m *Mock) GetRule(*nftables.Table, *nftables.Chain) ([]*nftables.Rule, error) {
	return nil, nil
}

// GetSetElements not implemented yet
func (m *Mock) GetSetElements(*nftables.Set) ([]nftables.SetElement, error) {
	return nil, nil
}

// GetSets not implemented yet
func (m *Mock) GetSets(*nftables.Table) ([]*nftables.Set, error) {
	return nil, nil
}

// ListChains not implemented yet
func (m *Mock) ListChains() ([]*nftables.Chain, error) {
	return nil, nil
}

// ListTables not implemented yet
func (m *Mock) ListTables() ([]*nftables.Table, error) {
	return nil, nil
}

// InitMockConn initializes mock connection of the nftables family
func InitMockConn() *Mock {
	m := &Mock{}
	m.ti = nftableslib.InitNFTables(m)
	return m
}
