func TestOrdinal(t *testing.T) {
	ord := ordinal(1)
	exp := "1st"
	if ord != exp {
		t.Error("expected %s, got %s", exp, ord)
	}
	
	ord = ordinal(10)
	exp = "10th"
	if ord != exp {
		t.Error("expected %s, got %s", exp, ord)
	}
}