import {graphql, useLazyLoadQuery} from "react-relay";
import {IndexPageQuery} from "relay/__generated__/IndexPageQuery.graphql";


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
      <div>hello</div>
    </div>
  )
}

