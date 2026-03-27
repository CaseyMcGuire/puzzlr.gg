import {graphql, useLazyLoadQuery} from "react-relay";
import {IndexPageQuery} from "relay/__generated__/IndexPageQuery.graphql";
import * as stylex from "@stylexjs/stylex";
import {create} from "@stylexjs/stylex";
import SidebarPageWrapper from "components/SidebarPageWrapper";

const styles = create({
  body: {
    backgroundColor: 'green'
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
    <SidebarPageWrapper>
      {
        query.users.map(users => {
          return (
            <div>
              {users.games?.map(game => game.id ?? "asdf")}
            </div>
          )
        })
      }
      <div sx={styles.body}>hello</div>
    </SidebarPageWrapper>
  )
}
