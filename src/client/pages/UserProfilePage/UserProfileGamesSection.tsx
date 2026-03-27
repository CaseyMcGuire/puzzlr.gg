import SectionCard from "components/SectionCard";
import {UserProfileSectionStyles as styles} from "pages/UserProfilePage/UserProfilePage.styles";
import {graphql, useFragment} from "react-relay";
import {UserProfileGamesSection_user$key} from "relay/__generated__/UserProfileGamesSection_user.graphql";

type Props = {
  user: UserProfileGamesSection_user$key;
};

export default function UserProfileGamesSection({user}: Props) {
  const data = useFragment(graphql`
    fragment UserProfileGamesSection_user on User {
      id
      games {
        id
        type
        status
        winner {
          id
        }
        currentTurn {
          id
        }
        user {
          id
          email
        }
      }
    }
  `, user);
  const games = data.games ?? [];
  const userID = data.id;

  return (
    <SectionCard title="Games">
      {games.length === 0 ? (
        <p sx={styles.emptyState}>This user has not joined any games yet.</p>
      ) : (
        <div sx={styles.list}>
          {games.map(game => {
            const opponents = (game.user ?? []).filter(participant => participant.id !== userID);
            const opponentLabel = opponents.length === 0
              ? "Solo"
              : opponents.map(opponent => opponent.email).join(", ");

            return (
              <div key={game.id} sx={styles.listRow}>
                <p sx={styles.rowTitle}>Game #{game.id}</p>
                <div sx={styles.rowMeta}>
                  <span>{game.type}</span>
                  <span>{game.status}</span>
                  <span>Opponent: {opponentLabel}</span>
                  {game.winner?.id === userID ? <span>Result: Won</span> : null}
                  {game.currentTurn?.id === userID ? <span>Current turn</span> : null}
                </div>
              </div>
            );
          })}
        </div>
      )}
    </SectionCard>
  );
}
