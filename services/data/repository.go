package data

type repository struct {
	work *unitOfWork
}

func newRepository(work []*unitOfWork) *repository {
	if len(work) > 0 {
		return &repository{work: work[0]}
	}
	return &repository{&unitOfWork{db: db}}
}
