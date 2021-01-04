package ds_test

import (
	"fmt"
	"testing"

	"github.com/mm4tt/goutil/ds"
)

func TestBitSet_Contains(t *testing.T) {
	testcases := []struct {
		bs               ds.BitSet
		shouldContain    []int
		shouldNotContain []int
	}{
		{
			bs:               ds.CreateBitSet(1, 2, 3),
			shouldContain:    []int{1, 2, 3},
			shouldNotContain: []int{0, 5, 6},
		},
		{
			bs:               ds.CreateBitSet(),
			shouldNotContain: []int{0, 1, 2},
		},
		{
			bs:               ds.CreateBitSet(10),
			shouldContain:    []int{10},
			shouldNotContain: []int{0, 1, 9, 16, 31, 63},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("Testcases #%d", i), func(t *testing.T) {
			for _, i := range tc.shouldContain {
				if !tc.bs.Contains(i) {
					t.Errorf("BitSet %v should contain %d", tc.bs, i)
				}
			}
			for _, i := range tc.shouldNotContain {
				if tc.bs.Contains(i) {
					t.Errorf("BitSet %v should contain %d", tc.bs, i)
				}
			}
		})
	}
}

func TestBitSet_Delete(t *testing.T) {
	testcases := []struct {
		bs     ds.BitSet
		delete []int
		want   ds.BitSet
	}{
		{
			bs:     ds.CreateBitSet(1, 2, 3),
			delete: []int{1, 3},
			want:   ds.CreateBitSet(2),
		},
		{
			bs:     ds.CreateBitSet(1),
			delete: []int{2},
			want:   ds.CreateBitSet(1),
		},
		{
			bs:     ds.CreateBitSet(10),
			delete: []int{10},
			want:   ds.CreateBitSet(),
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("Testcase #%d", i), func(t *testing.T) {
			for _, i := range tc.delete {
				tc.bs.Delete(i)
			}
			if tc.bs != tc.want {
				t.Errorf("want = %v, got = %v", tc.want, tc.bs)
			}
		})
	}
}

func TestBitSet_Intersect(t *testing.T) {
	testcases := []struct {
		bs1, bs2 ds.BitSet
		want     ds.BitSet
	}{
		{
			bs1:  ds.CreateBitSet(1, 2, 3),
			bs2:  ds.CreateBitSet(1, 2, 3),
			want: ds.CreateBitSet(1, 2, 3),
		},
		{
			bs1:  ds.CreateBitSet(1, 2, 3),
			bs2:  ds.CreateBitSet(3, 4, 5),
			want: ds.CreateBitSet(3),
		},
		{
			bs1:  ds.CreateBitSet(1, 2, 3),
			bs2:  ds.CreateBitSet(4, 5, 6),
			want: ds.CreateBitSet(),
		},
		{
			bs1:  ds.CreateBitSet(1, 2, 3),
			bs2:  ds.CreateBitSet(),
			want: ds.CreateBitSet(),
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("Testcases #%d", i), func(t *testing.T) {
			b := tc.bs1
			b.Intersect(tc.bs2)
			if b != tc.want {
				t.Errorf("want = %v, got = %v", tc.want, b)
			}
		})
	}
}
