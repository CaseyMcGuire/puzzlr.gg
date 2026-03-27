/**
 * @generated SignedSource<<50aaddd95525243dc730d43730543c48>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ReaderFragment } from 'relay-runtime';
import { FragmentRefs } from "relay-runtime";
export type UserProfileFriendsSection_user$data = {
  readonly friends: ReadonlyArray<{
    readonly email: string;
    readonly id: string;
  }> | null | undefined;
  readonly " $fragmentType": "UserProfileFriendsSection_user";
};
export type UserProfileFriendsSection_user$key = {
  readonly " $data"?: UserProfileFriendsSection_user$data;
  readonly " $fragmentSpreads": FragmentRefs<"UserProfileFriendsSection_user">;
};

const node: ReaderFragment = {
  "argumentDefinitions": [],
  "kind": "Fragment",
  "metadata": null,
  "name": "UserProfileFriendsSection_user",
  "selections": [
    {
      "alias": null,
      "args": null,
      "concreteType": "User",
      "kind": "LinkedField",
      "name": "friends",
      "plural": true,
      "selections": [
        {
          "alias": null,
          "args": null,
          "kind": "ScalarField",
          "name": "id",
          "storageKey": null
        },
        {
          "alias": null,
          "args": null,
          "kind": "ScalarField",
          "name": "email",
          "storageKey": null
        }
      ],
      "storageKey": null
    }
  ],
  "type": "User",
  "abstractKey": null
};

(node as any).hash = "8d5e85d5ea020970f0632b1dd354f6f3";

export default node;
