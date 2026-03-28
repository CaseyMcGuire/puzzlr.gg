/**
 * @generated SignedSource<<380303659053790b2ee3bcd91fbfa0b7>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest } from 'relay-runtime';
export type SendFriendRequestInput = {
  recipientID: string;
};
export type UserProfileAddFriendButtonMutation$variables = {
  input: SendFriendRequestInput;
};
export type UserProfileAddFriendButtonMutation$data = {
  readonly sendFriendRequest: {
    readonly __typename: "SendFriendRequestError";
    readonly message: string;
  } | {
    readonly __typename: "SendFriendRequestSuccess";
    readonly recipient: {
      readonly id: string;
    };
  } | {
    // This will never be '%other', but we need some
    // value in case none of the concrete values match.
    readonly __typename: "%other";
  };
};
export type UserProfileAddFriendButtonMutation = {
  response: UserProfileAddFriendButtonMutation$data;
  variables: UserProfileAddFriendButtonMutation$variables;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "defaultValue": null,
    "kind": "LocalArgument",
    "name": "input"
  }
],
v1 = [
  {
    "alias": null,
    "args": [
      {
        "kind": "Variable",
        "name": "input",
        "variableName": "input"
      }
    ],
    "concreteType": null,
    "kind": "LinkedField",
    "name": "sendFriendRequest",
    "plural": false,
    "selections": [
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "__typename",
        "storageKey": null
      },
      {
        "kind": "InlineFragment",
        "selections": [
          {
            "alias": null,
            "args": null,
            "concreteType": "User",
            "kind": "LinkedField",
            "name": "recipient",
            "plural": false,
            "selections": [
              {
                "alias": null,
                "args": null,
                "kind": "ScalarField",
                "name": "id",
                "storageKey": null
              }
            ],
            "storageKey": null
          }
        ],
        "type": "SendFriendRequestSuccess",
        "abstractKey": null
      },
      {
        "kind": "InlineFragment",
        "selections": [
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "message",
            "storageKey": null
          }
        ],
        "type": "SendFriendRequestError",
        "abstractKey": null
      }
    ],
    "storageKey": null
  }
];
return {
  "fragment": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Fragment",
    "metadata": null,
    "name": "UserProfileAddFriendButtonMutation",
    "selections": (v1/*: any*/),
    "type": "Mutation",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "UserProfileAddFriendButtonMutation",
    "selections": (v1/*: any*/)
  },
  "params": {
    "cacheID": "67f18f6bcfbe0ffc31df1fd5a3e0b749",
    "id": null,
    "metadata": {},
    "name": "UserProfileAddFriendButtonMutation",
    "operationKind": "mutation",
    "text": "mutation UserProfileAddFriendButtonMutation(\n  $input: SendFriendRequestInput!\n) {\n  sendFriendRequest(input: $input) {\n    __typename\n    ... on SendFriendRequestSuccess {\n      recipient {\n        id\n      }\n    }\n    ... on SendFriendRequestError {\n      message\n    }\n  }\n}\n"
  }
};
})();

(node as any).hash = "1bab3ad6d28183c201e58dbf8177910d";

export default node;
