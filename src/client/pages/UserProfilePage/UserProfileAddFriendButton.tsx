import * as stylex from "@stylexjs/stylex";
import {useState} from "react";
import {graphql, useMutation} from "react-relay";
import {UserProfileAddFriendButtonMutation} from "relay/__generated__/UserProfileAddFriendButtonMutation.graphql";

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
  message: {
    margin: 0,
    fontSize: "0.9rem",
    color: "#475569",
    textAlign: "right",
  },
  messageError: {
    color: "#b91c1c",
  },
  messageSuccess: {
    color: "#166534",
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
};

type MessageState = {
  tone: "error" | "success";
  text: string;
} | null;

export default function UserProfileAddFriendButton({recipientID}: Props) {
  const [commit, isInFlight] = useMutation<UserProfileAddFriendButtonMutation>(mutation);
  const [message, setMessage] = useState<MessageState>(null);
  const [isSent, setIsSent] = useState(false);

  const isDisabled = isInFlight || isSent;

  const handleClick = () => {
    if (isDisabled) {
      return;
    }

    setMessage(null);

    commit({
      variables: {
        input: {
          recipientID,
        },
      },
      onCompleted: (response, errors) => {
        if (errors && errors.length > 0) {
          setMessage({
            tone: "error",
            text: "Unable to send friend request right now.",
          });
          return;
        }

        const result = response.sendFriendRequest;
        if (!result) {
          setMessage({
            tone: "error",
            text: "Unable to send friend request right now.",
          });
          return;
        }

        if (result.__typename === "SendFriendRequestError") {
          setMessage({
            tone: "error",
            text: result.message,
          });
          return;
        }

        if (result.__typename !== "SendFriendRequestSuccess") {
          setMessage({
            tone: "error",
            text: "Unable to send friend request right now.",
          });
          return;
        }

        setIsSent(true);
        setMessage({
          tone: "success",
          text: "Friend request sent!",
        });
      },
      onError: () => {
        setMessage({
          tone: "error",
          text: "Unable to send friend request right now.",
        });
      },
    });
  };

  return (
    <div sx={styles.root}>
      <button
        type="button"
        sx={[styles.button, isDisabled && styles.buttonDisabled]}
        disabled={isDisabled}
        onClick={handleClick}
      >
        {isInFlight ? "Sending..." : isSent ? "Request Sent" : "Add Friend"}
      </button>
      {message ? (
        <p
          sx={[
            styles.message,
            message.tone === "error" && styles.messageError,
            message.tone === "success" && styles.messageSuccess,
          ]}
        >
          {message.text}
        </p>
      ) : null}
    </div>
  );
}
