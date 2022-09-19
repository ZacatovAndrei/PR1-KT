package main

type Cook struct {
	id, rank, proficiency int
}

func (c *Cook) Init(id, rank, proficiency int) {
	c.id = id
	c.rank = rank
	c.proficiency = proficiency
}
