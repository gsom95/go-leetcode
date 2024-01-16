package maxconsecutiveones

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func Test_findMaxConsecutiveOnes(t *testing.T) {
	paths, err := filepath.Glob(filepath.Join("testdata", "*.txt"))
	if err != nil {
		t.Fatal(err)
	}

	for _, path := range paths {
		_, filename := filepath.Split(path)
		testname := filename[:len(filename)-len(filepath.Ext(path))]

		t.Run(testname, func(t *testing.T) {
			source, err := os.Open(path)
			if err != nil {
				t.Fatal("error opening source file:", err)
			}
			scanner := bufio.NewScanner(source)
			scanner.Split(bufio.ScanLines)

			scanner.Scan()
			inputStr := strings.Split(scanner.Text(), ",")
			input := func() []int {
				res := make([]int, len(inputStr))
				for i, str := range inputStr {
					res[i], err = strconv.Atoi(str)
					if err != nil {
						t.Fatal(err)
					}
				}
				return res
			}()

			scanner.Scan()
			want, _ := strconv.Atoi(scanner.Text())

			if got := findMaxConsecutiveOnes(input); got != want {
				t.Errorf("findMaxConsecutiveOnes() = %v, want %v", got, want)
			}
		})
	}
}
