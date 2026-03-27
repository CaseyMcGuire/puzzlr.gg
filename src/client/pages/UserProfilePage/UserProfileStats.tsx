import KpiCard from "components/KpiCard";
import * as stylex from "@stylexjs/stylex";
import {graphql, useFragment} from "react-relay";
import {UserProfileStats_user$key} from "relay/__generated__/UserProfileStats_user.graphql";

const styles = stylex.create({
  statsGrid: {
    display: "flex",
    flexWrap: "wrap",
    gap: "14px",
    marginBottom: "24px",
  },
});

type Props = {
  user: UserProfileStats_user$key;
};

export default function UserProfileStats({user}: Props) {
  const data = useFragment(graphql`
    fragment UserProfileStats_user on User {
      id
      friends {
        id
      }
      games {
        winner {
          id
        }
      }
    }
  `, user);
  const friends = data.friends ?? [];
  const games = data.games ?? [];
  const wins = games.filter(game => game.winner?.id === data.id).length;

  return (
    <div sx={styles.statsGrid}>
      <KpiCard label="Friends" value={friends.length} />
      <KpiCard label="Games" value={games.length} />
      <KpiCard label="Wins" value={wins} />
    </div>
  );
}
