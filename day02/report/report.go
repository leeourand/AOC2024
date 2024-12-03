package report

type Level int
type Report []Level

func (r Report) IsSafeWithTolerances() bool {
	if r.hasError() {
		for _, mutation := range r.mutations() {
			if !mutation.hasError() {
				return true
			}
		}
	} else {
		return true
	}
	return false
}

func (r Report) mutations() []Report {
	var mutations []Report

	for i := range len(r) {
		mutation := append(Report([]Level{}), r[:i]...)
		mutation = append(mutation, r[i+1:]...)
		mutations = append(mutations, mutation)
	}

	return mutations
}

func (r Report) IsSafe() bool {
	return !r.hasError()
}

func (r Report) hasError() bool {
	var ascending bool

	for i, level := range r {
		if i > 0 {
			previous := r[i-1]
			if ascending {
				if level <= previous {
					return true
				} else if (level - previous) > 3 {
					return true
				}
			} else {
				if level >= previous {
					return true
				} else if (previous - level) > 3 {
					return true
				}
			}
		} else {
			ascending = level < r[i+1]
		}
	}
	return false
}
