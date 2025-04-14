import {graphql, useLazyLoadQuery} from "react-relay";
import {IndexPageQuery} from "relay/__generated__/IndexPageQuery.graphql";


export default function IndexPage() {
  const query = useLazyLoadQuery<IndexPageQuery>(graphql`
    query IndexPageQuery {
        todos {
            id
        }
    }
  `, {});

  return (
    <div>
      {
        query.todos.map(todo => {
          return (
            <div>
              {todo.id}
            </div>
          )
        })
      }
      <div>hello</div>
    </div>
  )
}

