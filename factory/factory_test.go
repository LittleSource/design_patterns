package factory

import "testing"

func TestNewSimpleFactory(t *testing.T) {
	factory := NewSimpleFactory()
	family := factory.GetFamilyFood()
	family.Cook()
	school := factory.GetSchoolFood()
	school.Cook()
}
