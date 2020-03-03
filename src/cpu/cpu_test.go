package cpu

import "testing"

func TestCPUObjectExists(t *testing.T) {
	var c *CPU
	c = Make()
	if c == nil {
		t.Fatal("CPU Object is null")
	}
}

func TestCPUHasInitializedFlags(t *testing.T) {
	c := Make()
	if c.x != 0 {
		t.Fatal("CPU should have X register initalized to 0")
	}

	if c.y != 0 {
		t.Fatal("CPU should have Y register initalized to 0")
	}

	if c.a != 0 {
		t.Fatal("CPU should have Accumulator initialized to 0")
	}

	if c.s != 0 {
		t.Fatal("CPU should have Stack Pointer initialized to 0")
	}

	if c.p != 0 {
		t.Fatal("CPU should have flags register initialzed to 0")
	}

	if c.pc != 0 {
		t.Fatal("CPU should have Program Counter initialzed to 0")
	}

	if c.irq != 0 {
		t.Fatal("CPU should have IRQ initialized to 0")
	}
}

func TestLoadImmediateAccumulator(t *testing.T) {
	c := Make()
	c.LoadImmediateA(5)
	if c.a != 5 {
		t.Fatal("Accumulator should have value 5")
	}
}

func TestLoadImmediateX(t *testing.T) {
	c := Make()
	c.LoadImmediateX(7)
	if c.x != 7 {
		t.Fatal("Register X should have value 7")
	}
}

func TestLoadImmediateY(t *testing.T) {
	c := Make()
	c.LoadImmediateY(10)
	if c.y != 10 {
		t.Fatal("Register Y should have value 10")
	}
}
