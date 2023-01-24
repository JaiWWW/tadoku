-- name: LeaderboardForContest :many
with leaderboard as (
  select
    user_id,
    sum(score) as score
  from logs
  inner join contest_logs
    on contest_logs.log_id = logs.id
  where
    contest_logs.contest_id = sqlc.arg('contest_id')
    and logs.deleted_at is null
    and (logs.language_code = sqlc.narg('language_code') or sqlc.narg('language_code') is null)
    and (logs.log_activity_id = sqlc.narg('activity_id')::integer or sqlc.narg('activity_id') is null)
  group by user_id
), ranked_leaderboard as (
  select
    user_id,
    score,
    rank() over(order by score desc) as "rank"
  from leaderboard
), registrations as (
  select
    id,
    user_id,
    user_display_name
  from contest_registrations
  where
    contest_id = sqlc.arg('contest_id')
    and deleted_at is null
    and (sqlc.narg('language_code') = any(language_codes) or sqlc.narg('language_code') is null)
)
select
  rank() over(order by score desc) as "rank",
  registrations.user_id,
  registrations.user_display_name,
  coalesce(ranked_leaderboard.score, 0)::real as score,
  coalesce((
    "rank" = lag("rank", 1, -1::bigint) over (order by "rank")
    or "rank" = lead("rank", 1, -1::bigint) over (order by "rank")
  ), false)::boolean as is_tie,
  (select count(registrations.user_id) from registrations) as total_size
from registrations
left join ranked_leaderboard using(user_id)
order by
  score desc,
  registrations.user_id asc
limit sqlc.arg('page_size')
offset sqlc.arg('start_from');

-- name: OfficialLeaderboardPreviewForYear :many
with leaderboard as (
  select
    user_id,
    sum(score) as score
  from logs
  inner join contest_logs
    on contest_logs.log_id = logs.id
  where
    logs.year = sqlc.arg('year')
    and eligible_official_leaderboard = true
    and logs.deleted_at is null
  group by user_id
), ranked_leaderboard as (
  select
    user_id,
    score,
    rank() over(order by score desc) as "rank"
  from leaderboard
), registrations as (
  select
    id,
    user_id,
    user_display_name
  from contest_registrations
  where
    extract(year from created_at) = sqlc.arg('year')::integer
    and deleted_at is null
)
select
  rank() over(order by score desc) as "rank",
  registrations.user_id,
  registrations.user_display_name,
  coalesce(ranked_leaderboard.score, 0)::real as score,
  coalesce((
    "rank" = lag("rank", 1, -1::bigint) over (order by "rank")
    or "rank" = lead("rank", 1, -1::bigint) over (order by "rank")
  ), false)::boolean as is_tie
from registrations
left join ranked_leaderboard using(user_id)
order by
  score desc,
  registrations.user_id asc
limit 10;