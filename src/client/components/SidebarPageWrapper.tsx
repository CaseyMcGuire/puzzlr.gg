import Sidebar from "components/Sidebar/Sidebar";
import {ReactNode} from "react";
import {SidebarStyles} from "components/Sidebar/SidebarStyles.stylex";
import * as stylex from "@stylexjs/stylex";

type Props = {
  children: ReactNode
}

const styles = stylex.create({
  wrapper: {
    display: 'flex',
    flexDirection: 'row'
  },
  sidebarPageContents: {
    marginLeft: SidebarStyles.sidebarWidth,
    height: '100%',
    width: '100%',
  }
})

export default function SidebarPageWrapper(props: Props) {
  return (
    <div sx={styles.wrapper}>
      <Sidebar items={[
        {type: 'link', label: 'Home', href: '/'},
        {type: 'link', label: 'Tic-Tac-Toe', href: '/tictactoe'},
      ]} />
      <div sx={styles.sidebarPageContents}>
        {props.children}
      </div>
    </div>
  )
}
