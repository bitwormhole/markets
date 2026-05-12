package normalizers

import (
	"testing"
)

func TestNormalizeRegularName(t *testing.T) {

	list := make([]string, 0)

	list = append(list, "apple")
	list = append(list, "balabala (ltd) xyz")
	list = append(list, "xxx（曹县）有限公司")

	for index, n1 := range list {

		n2 := NormalizeRegularName(n1)

		t.Logf("\n name[%d]: \n  n1 = %s \n  n2 = %s \n", index, n1, n2)

	}

}
