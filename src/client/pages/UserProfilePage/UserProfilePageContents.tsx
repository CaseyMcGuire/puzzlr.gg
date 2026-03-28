import * as stylex from "@stylexjs/stylex";
import {graphql, useLazyLoadQuery} from "react-relay";
import UserProfileAddFriendButton from "pages/UserProfilePage/UserProfileAddFriendButton";
import UserProfileFriendsSection from "pages/UserProfilePage/UserProfileFriendsSection";
import UserProfileGamesSection from "pages/UserProfilePage/UserProfileGamesSection";
import UserProfileHero from "pages/UserProfilePage/UserProfileHero";
import UserProfileStats from "pages/UserProfilePage/UserProfileStats";
import {UserProfilePageContentsQuery} from "relay/__generated__/UserProfilePageContentsQuery.graphql";

export const userProfilePageStyles = stylex.create({
  page: {
    minHeight: "100%",
    padding: "32px",
    backgroundColor: "#ffffff",
    color: "#111827",
    fontFamily: "Segoe UI, Tahoma, Geneva, Verdana, sans-serif",
  },
  sectionGrid: {
    display: "flex",
    flexWrap: "wrap",
    gap: "18px",
  },
});

type Props = {
  userID: number;
};

export default function UserProfilePageContents({userID}: Props) {
  const query = useLazyLoadQuery<UserProfilePageContentsQuery>(graphql`
    query UserProfilePageContentsQuery($id: ID!) {
      user(id: $id) {
        id
        email
        viewerFriendshipStatus
        ...UserProfileFriendsSection_user
        ...UserProfileStats_user
        ...UserProfileGamesSection_user
      }
    }
  `, {
    id: String(userID),
  });

  const user = query.user;

  if (!user) {
    return (
      <div sx={userProfilePageStyles.page}>
        <UserProfileHero
          title="User not found"
          subtitle={`No user exists with ID ${userID}.`}
        />
      </div>
    );
  }

  return (
    <div sx={userProfilePageStyles.page}>
      <UserProfileHero
        title={user.email}
        subtitle={`User ID ${user.id}`}
        actions={
          <UserProfileAddFriendButton
            recipientID={user.id}
            viewerFriendshipStatus={user.viewerFriendshipStatus}
          />
        }
      />
      <UserProfileStats user={user} />

      <div sx={userProfilePageStyles.sectionGrid}>
        <UserProfileFriendsSection user={user} />
        <UserProfileGamesSection user={user} />
      </div>
    </div>
  );
}
