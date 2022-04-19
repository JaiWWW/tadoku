package repositories

import (
	"github.com/tadoku/tadoku/services/reading-contest-api/domain"
	"github.com/tadoku/tadoku/services/reading-contest-api/interfaces/rdb"
	"github.com/tadoku/tadoku/services/reading-contest-api/usecases"
)

// NewContestRepository instantiates a new contest repository
func NewContestRepository(sqlHandler rdb.SQLHandler) usecases.ContestRepository {
	return &contestRepository{sqlHandler: sqlHandler}
}

type contestRepository struct {
	sqlHandler rdb.SQLHandler
}

func (r *contestRepository) Store(contest *domain.Contest) error {
	if contest.ID == 0 {
		return r.create(contest)
	}

	return r.update(contest)
}

func (r *contestRepository) create(contest *domain.Contest) error {
	query := `
		insert into contests
		(description, start, "end", open)
		values ($1, $2, $3, $4)
		returning id
	`

	row := r.sqlHandler.QueryRow(query, contest.Description, contest.Start, contest.End, contest.Open)
	err := row.Scan(&contest.ID)
	if err != nil {
		return domain.WrapError(err)
	}

	return nil
}

func (r *contestRepository) update(contest *domain.Contest) error {
	query := `
		update contests
		set description = :description, start = :start, "end" = :end, open = :open
		where id = :id
	`

	_, err := r.sqlHandler.NamedExecute(query, contest)
	return domain.WrapError(err)
}

func (r *contestRepository) GetOpenContests() ([]uint64, error) {
	query := `
		select id
		from contests
		where open = true
	`

	var ids []uint64
	err := r.sqlHandler.Select(&ids, query)
	if err != nil {
		return nil, domain.WrapError(err)
	}

	return ids, nil
}

func (r *contestRepository) GetRunningContests() ([]uint64, error) {
	query := `
		select id
		from contests
		where
			start <= now() at time zone 'utc' and
			"end" >= now() at time zone 'utc'
	`

	var ids []uint64
	err := r.sqlHandler.Select(&ids, query)
	if err != nil {
		return nil, domain.WrapError(err)
	}

	return ids, nil
}

func (r *contestRepository) FindAll() ([]domain.Contest, error) {
	query := `
		select id, description, start, "end", open
		from contests
		order by id desc
	`

	var contests []domain.Contest
	err := r.sqlHandler.Select(&contests, query)
	if err != nil {
		return contests, domain.WrapError(err)
	}

	if len(contests) == 0 {
		return nil, domain.ErrNotFound
	}

	return contests, nil
}

func (r *contestRepository) FindRecent(count int) ([]domain.Contest, error) {
	query := `
		select id, description, start, "end", open
		from contests
		order by id desc
		limit $1
	`

	var contests []domain.Contest
	err := r.sqlHandler.Select(&contests, query, count)
	if err != nil {
		return contests, domain.WrapError(err)
	}

	return contests, nil
}

func (r *contestRepository) FindByID(id uint64) (domain.Contest, error) {
	query := `
		select id, description, start, "end", open
		from contests
		where id = $1
		limit 1
	`

	var contest domain.Contest
	err := r.sqlHandler.Get(&contest, query, id)
	if err != nil {
		return contest, domain.WrapError(err)
	}

	return contest, nil
}

func (r *contestRepository) Stats(id uint64) (domain.ContestStats, error) {
	var contestStats domain.ContestStats = domain.ContestStats{ByLanguage: []domain.ContestLanguageStat{}, Participants: 0, TotalAmount: 0}

	_, err := r.FindByID(id)

	if err != nil {
		return contestStats, err
	}

	queryStatsByLanguage := `
		select count(*) as cnt, language_code
		from rankings
		where
			amount > 0 and
			contest_id = $1
		group by
			language_code
		order by cnt desc
	`
	queryParticipants := `
		select count(*)
			from rankings
		where
			language_code = 'GLO' and
			amount > 0 and
			contest_id = $1
	`

	queryTotalPages := `
		select sum(amount)
		from rankings
		where
			language_code = 'GLO' and
			amount > 0 and
			contest_id = $1
	`

	err = r.sqlHandler.Select(&contestStats.ByLanguage, queryStatsByLanguage, id)
	if err != nil {
		return contestStats, domain.WrapError(err)
	}

	if len(contestStats.ByLanguage) == 0 {
		return contestStats, nil
	}

	err = r.sqlHandler.Get(&contestStats.Participants, queryParticipants, id)
	if err != nil {
		return contestStats, domain.WrapError(err)
	}

	err = r.sqlHandler.Get(&contestStats.TotalAmount, queryTotalPages, id)
	if err != nil {
		return contestStats, domain.WrapError(err)
	}

	return contestStats, nil
}
