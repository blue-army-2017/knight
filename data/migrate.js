import { Database } from "bun:sqlite";

const dataPath = Bun.argv[2];
const dbPath = Bun.argv[3];

const data = await Bun.file(dataPath).json();
const db = new Database(dbPath);

try {
  const insertMember = db.prepare(`INSERT OR REPLACE INTO members
(id, first_name, last_name, active)
values
($id, $firstName, $lastName, $active)`);
  const insertMembers = db.transaction((members) => {
    for (const [memberId, member] of Object.entries(members)) {
      insertMember.run({
        $id: memberId,
        $firstName: member.firstName,
        $lastName: member.lastName,
        $active: member.active ? 1 : 0,
      });
    }
  });

  const insertSeason = db.prepare(`INSERT OR REPLACE INTO seasons
(id, name, created)
values
($id, $name, $created)`);
  const insertSeasons = db.transaction((seasons) => {
    for (const [seasonId, season] of Object.entries(seasons)) {
      const created = new Date(season.created);
      insertSeason.run({
        $id: seasonId,
        $name: season.name,
        $created: getDateRepresentation(created),
      });
    }
  });

  const insertSeasonGame = db.prepare(`INSERT OR REPLACE INTO season_games
(id, opponent, home, mode, date, season_id)
values
($id, $opponent, $home, $mode, $date, $season_id)`);
  const insertSeasonGames = db.transaction((season) => {
    for (const [gameId, game] of Object.entries(season.games)) {
      const date = new Date(game.date);
      insertSeasonGame.run({
        $id: gameId,
        $opponent: game.opponent,
        $home: game.home ? 1 : 0,
        $mode: game.mode,
        $date: getDateRepresentation(date),
        $season_id: season.id,
      });
    }
  });

  const insertGamePresence = db.prepare(`INSERT OR REPLACE INTO present_members
(game_id, member_id)
values
($game_id, $member_id)`);
  const insertGamePresences = db.transaction((game) => {
    for (const [_, memberId] of Object.entries(game.presentMembers)) {
      insertGamePresence.run({
        $game_id: game.id,
        $member_id: memberId,
      });
    }
  });

  insertMembers(data.member);
  insertSeasons(data.season);
  for (const [seasonId, season] of Object.entries(data.season)) {
    insertSeasonGames({ id: seasonId, games: season.games || {} });

    for (const [gameId, game] of Object.entries(season.games || {})) {
      insertGamePresences({ id: gameId, presentMembers: game.presentMembers });
    }
  }
} finally {
  db.close();
}

function getDateRepresentation(date) {
  const day = date.getUTCDate().toString().padStart(2, "0");
  const month = (date.getUTCMonth() + 1).toString().padStart(2, "0");
  const year = date.getUTCFullYear();

  return `${year}-${month}-${day}`;
}
