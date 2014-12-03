package mem

import "testing"

func TestPhysicalMem(t *testing.T) {

	mem, err := PhysicalMem()
	if err != nil {
		t.Errorf("Unable to get physical memory: %v", err)
	}
	if mem <= 0 {
		t.Errorf("Unexpected memory value!")
	}
	t.Logf("Physical memory is %d", mem)
}
