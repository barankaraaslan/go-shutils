package shutil

import (
	"fmt"
	"os"
	"testing"
)

// cp src dst
func TestCopyAFile(t *testing.T) {
	// Clean up existing test files from previous attemts
	if err := os.RemoveAll("src"); err != nil { t.Fatal(err) }
	if err := os.RemoveAll("dst"); err != nil { t.Fatal(err) }
	
	// Create test files
	if _, err := os.Create("src"); err != nil { t.Fatal(err) }

	if err := Copy("src", "dst"); err != nil { t.Fatal(err) }
	if _, err := os.Stat("dst"); err != nil {
		if os.IsNotExist(err) {
			t.Fail()
		} else {
			t.Fatal(err)
		}
	}

	// Clean up
	if err := os.Remove("src"); err != nil { t.Fatal(err) }
	if err := os.Remove("dst"); err != nil { t.Fatal(err) }
}

// cp bin/../bin/src bin/../dst
func TestCopyAFileWithRelativePaths(t *testing.T) {
	// Clean up existing test files from previous attemts
	if err := os.RemoveAll("bin"); err != nil { t.Fatal(err) }
	if err := os.RemoveAll("dst"); err != nil { t.Fatal(err) }
	
	// Create test files
	if err := os.Mkdir("bin", 0755); err != nil { t.Fatal(err) }
	if _, err := os.Create("bin/src"); err != nil { t.Fatal(err) }

	if err := Copy("bin/../bin/src", "bin/../dst"); err != nil { t.Fatal(err) }
	if _, err := os.Stat("dst"); err != nil {
		if os.IsNotExist(err) {
			t.Fail()
		} else {
			t.Fatal(err)
		}
	}

	// Clean up
	if err := os.RemoveAll("bin"); err != nil { t.Fatal(err) }
	if err := os.Remove("dst"); err != nil { t.Fatal(err) }
}

// cp src bin/dst
func TestCopyAFileIntoExistingDir(t *testing.T) {
	// Clean up existing test files from previous attemts
	if err := os.RemoveAll("src"); err != nil { t.Fatal(err) }
	if err := os.RemoveAll("bin"); err != nil { t.Fatal(err) }
	
	// Create test files
	if _, err := os.Create("src"); err != nil { t.Fatal(err) }
	if err := os.Mkdir("bin", 0755); err != nil { t.Fatal(err) }

	if err := Copy("src", "bin/dst"); err != nil { t.Fatal(err) }
	if _, err := os.Stat("bin/dst"); err != nil {
		if os.IsNotExist(err) {
			t.Fail()
		} else {
			t.Fatal(err)
		}
	}

	// Clean up
	if err := os.Remove("src"); err != nil { t.Fatal(err) }
	if err := os.RemoveAll("bin"); err != nil { t.Fatal(err) }
}

// cp src bin/
func TestCopyAFileIntoExistingDirWithoutDstName(t *testing.T) {
	// Clean up existing test files from previous attemts
	if err := os.RemoveAll("src"); err != nil { t.Fatal(err) }
	if err := os.RemoveAll("bin"); err != nil { t.Fatal(err) }
	
	// Create test files
	if _, err := os.Create("src"); err != nil { t.Fatal(err) }
	if err := os.Mkdir("bin", 0755); err != nil { t.Fatal(err) }

	if err := Copy("src", "bin/"); err != nil { t.Fatal(err) }
	if _, err := os.Stat("bin/src"); err != nil {
		if os.IsNotExist(err) {
			t.Fail()
		} else {
			t.Fatal(err)
		}
	}

	// Clean up
	if err := os.Remove("src"); err != nil { t.Fatal(err) }
	if err := os.RemoveAll("bin"); err != nil { t.Fatal(err) }
}

// cp src $PWD/bin/
func TestCopyAFileIntoExistingDirWithoutDstNameAndAbsDst(t *testing.T) {
	// Clean up existing test files from previous attemts
	if err := os.RemoveAll("src"); err != nil { t.Fatal(err) }
	if err := os.RemoveAll("bin"); err != nil { t.Fatal(err) }
	
	// Create test files
	if _, err := os.Create("src"); err != nil { t.Fatal(err) }
	if err := os.Mkdir("bin", 0755); err != nil { t.Fatal(err) }

	cwd, err := os.Getwd()
	if err != nil { t.Fatal(err) }

	if err := Copy("src", fmt.Sprintf("%s/bin/", cwd)); err != nil { t.Fatal(err) }
	if _, err := os.Stat("bin/src"); err != nil {
		if os.IsNotExist(err) {
			t.Fail()
		} else {
			t.Fatal(err)
		}
	}

	// Clean up
	if err := os.Remove("src"); err != nil { t.Fatal(err) }
	if err := os.RemoveAll("bin"); err != nil { t.Fatal(err) }
}

// cp dir1 dir2
func TestCopyADirectory(t *testing.T) {
	// Clean up existing test files from previous attemts
	if err := os.RemoveAll("dir1"); err != nil { t.Fatal(err) }
	if err := os.RemoveAll("dir2"); err != nil { t.Fatal(err) }
	
	// Create test files
	if err := os.Mkdir("dir1", 0755); err != nil { t.Fatal(err) }
	if err := os.Mkdir("dir2", 0755); err != nil { t.Fatal(err) }
	if _, err := os.Create("dir1/file1.txt"); err != nil { t.Fatal(err) }
	if _, err := os.Create("dir2/file2.txt"); err != nil { t.Fatal(err) }

	if err := Copy("dir1", "dir2"); err != nil { t.Fatal(err) }
	if _, err := os.Stat("dir2/dir1/file1.txt"); err != nil { if os.IsNotExist(err) { t.Fail() } else { t.Fatal(err) } }
	if _, err := os.Stat("dir2/file2.txt"); err != nil { if os.IsNotExist(err) { t.Fail() } else { t.Fatal(err) } }

	// Clean up
	if err := os.RemoveAll("dir1"); err != nil { t.Fatal(err) }
	if err := os.RemoveAll("dir2"); err != nil { t.Fatal(err) }
}