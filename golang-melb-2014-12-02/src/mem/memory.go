package mem

// PhysicalMem returns the machine's physical memory in bytes.
func PhysicalMem() (mem uint64, err error) {
	return physicalMem()
}
