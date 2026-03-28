import * as stylex from "@stylexjs/stylex";
import {useState} from "react";
import {graphql, useMutation} from "react-relay";
import {UserProfileAddFriendButtonMutation} from "relay/__generated__/UserProfileAddFriendButtonMutation.graphql";
import {ViewerFriendshipStatus} from "relay/__generated__/UserProfilePageContentsQuery.graphql";

const styles = stylex.create({
  root: {
    display: "flex",
    flexDirection: "column",
    alignItems: "flex-end",
    gap: "8px",
  },
  button: {
    borderWidth: 0,
    borderRadius: "999px",
    backgroundColor: "#1d4ed8",
    color: "#ffffff",
    padding: "10px 18px",
    fontSize: "0.95rem",
    fontWeight: "700",
    cursor: "pointer",
    transitionDuration: "150ms",
  },
  buttonDisabled: {
    backgroundColor: "#94a3b8",
    cursor: "default",
  },
});

const mutation = graphql`
  mutation UserProfileAddFriendButtonMutation($input: SendFriendRequestInput!) {
    sendFriendRequest(input: $input) {
      __typename
      ... on SendFriendRequestSuccess {
        recipient {
          id
        }
      }
      ... on SendFriendRequestError {
        message
      }
    }
  }
`;

type Props = {
  recipientID: string;
  viewerFriendshipStatus: ViewerFriendshipStatus;
};

export default function UserProfileAddFriendButton({recipientID, viewerFriendshipStatus}: Props) {
  switch (viewerFriendshipStatus) {
    case "NOT_FRIENDS":
    case "REQUEST_SENT":
    case "REQUEST_RECEIVED":
      return (
        <UserProfileAddFriendButtonImpl
          recipientID={recipientID}
          initialStatus={viewerFriendshipStatus}
        />
      );
    case "FRIENDS":
    case "NOT_LOGGED_IN":
    case "NOT_APPLICABLE":
    default:
      return null;
  }
}

type ImplProps = {
  recipientID: string;
  initialStatus: "NOT_FRIENDS" | "REQUEST_SENT" | "REQUEST_RECEIVED";
};

function UserProfileAddFriendButtonImpl({recipientID, initialStatus}: ImplProps) {
  const [commit, isInFlight] = useMutation<UserProfileAddFriendButtonMutation>(mutation);
  const [status, setStatus] = useState(initialStatus);

  const canSendRequest = status === "NOT_FRIENDS";
  const isDisabled = isInFlight || !canSendRequest;

  const handleClick = () => {
    if (isDisabled) {
      return;
    }

    commit({
      variables: {
        input: {
          recipientID,
        },
      },
      onCompleted: (response, errors) => {
        const result = response.sendFriendRequest;
        if ((errors && errors.length > 0) || !result) {
          return;
        }

        switch (result.__typename) {
          case "SendFriendRequestSuccess":
            setStatus("REQUEST_SENT");
            break;
        }
      },
    });
  };

  const buttonLabel = (() => {
    if (isInFlight) return "Sending...";
    switch (status) {
      case "REQUEST_SENT": return "Request Sent";
      case "REQUEST_RECEIVED": return "Accept Request";
      case "NOT_FRIENDS": return "Add Friend";
    }
  })();

  return (
    <div sx={styles.root}>
      <button
        type="button"
        sx={[styles.button, isDisabled && styles.buttonDisabled]}
        disabled={isDisabled}
        onClick={handleClick}
      >
        {buttonLabel}
      </button>
    </div>
  );
}
