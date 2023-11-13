package nomad

type Column struct {
	attr *attribute
}

func (c *Column) PrimaryKey() {
	c.attr.primaryKey = true
}

func (c *Column) Nullable() {
	c.attr.nullable = true
}
