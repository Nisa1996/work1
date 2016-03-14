package Calculate
import(
  "github.com/stretchr/testify/assert"
  "testing"

)

func TestSumFunc(t *testing.T){
  assert.Equal(t,SumFunc(2,3),5)
  assert.Equal(t,SumFunc(4,3),7)
}

func TestDifFunc(t *testing.T){
  assert.Equal(t,DifFunc(5,3),2)
  assert.Equal(t,DifFunc(6,3),3)
}

func TestMultiFunc(t *testing.T){
  assert.Equal(t,MultiFunc(2,3),6)
  assert.Equal(t,MultiFunc(3,3),9)
}

func TestDivFunc(t *testing.T){
  assert.Equal(t,DivFunc(6,3),2)
  assert.Equal(t,DivFunc(3,3),1)
}
