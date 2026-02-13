import {graphql, useLazyLoadQuery} from "react-relay";
import {IndexPageQuery} from "relay/__generated__/IndexPageQuery.graphql";
import * as stylex from "@stylexjs/stylex";

const styles = stylex.create({
  body: {
    backgroundColor: 'blue'
  }
});

export default function IndexPage() {
  const query = useLazyLoadQuery<IndexPageQuery>(graphql`
    query IndexPageQuery {
        users {
            games {
                id
            }
        }
    }
  `, {});

  return (
    <div>
      {
        query.users.map(users => {
          return (
            <div>
              {users.games?.map(game => game.id ?? "asdf")}
            </div>
          )
        })
      }
      <div {...stylex.props(styles.body)}>hello</div>
    </div>
  )
}

