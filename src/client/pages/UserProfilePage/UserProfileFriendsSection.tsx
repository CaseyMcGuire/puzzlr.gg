import SectionCard from "components/SectionCard";
import {UserProfileSectionStyles as styles} from "pages/UserProfilePage/UserProfilePage.styles";
import {graphql, useFragment} from "react-relay";
import {UserProfileFriendsSection_user$key} from "relay/__generated__/UserProfileFriendsSection_user.graphql";

type Props = {
  user: UserProfileFriendsSection_user$key;
};

export default function UserProfileFriendsSection({user}: Props) {
  const data = useFragment(graphql`
    fragment UserProfileFriendsSection_user on User {
      friends {
        id
        email
      }
    }
  `, user);
  const friends = data.friends ?? [];

  return (
    <SectionCard title="Friends">
      {friends.length === 0 ? (
        <p sx={styles.emptyState}>This user has no friends yet.</p>
      ) : (
        <div sx={styles.list}>
          {friends.map(friend => (
            <div key={friend.id} sx={styles.listRow}>
              <p sx={styles.rowTitle}>{friend.email}</p>
              <div sx={styles.rowMeta}>
                <span>ID {friend.id}</span>
                <a sx={styles.link} href={`/user/${friend.id}`}>
                  View profile
                </a>
              </div>
            </div>
          ))}
        </div>
      )}
    </SectionCard>
  );
}
