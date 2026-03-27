import SidebarPageWrapper from "components/SidebarPageWrapper";
import {useParams} from "react-router";
import UserProfileHero from "pages/UserProfilePage/UserProfileHero";
import UserProfilePageContents, {
  userProfilePageStyles,
} from "pages/UserProfilePage/UserProfilePageContents";

export default function UserProfilePage() {
  const params = useParams();
  const userID = Number(params.id);
  const isValidUserID = Number.isInteger(userID) && userID > 0;

  return (
    <SidebarPageWrapper>
      {isValidUserID ? (
        <UserProfilePageContents userID={userID} />
      ) : (
        <div sx={userProfilePageStyles.page}>
          <UserProfileHero
            title="Invalid user ID"
            subtitle="The URL must use a positive numeric user ID."
          />
        </div>
      )}
    </SidebarPageWrapper>
  );
}
