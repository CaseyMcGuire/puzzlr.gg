import Sidebar from "components/Sidebar/Sidebar";
import {ReactNode} from "react";
import {create} from "@stylexjs/stylex";
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
    marginLeft: SidebarStyles.sidebarWidth
  }
})

export default function SidebarPageWrapper(props: Props) {
  return (
    <div sx={styles.wrapper}>
      <Sidebar items={[
        {type: 'link', label: 'Home', href: '/'},
      ]} />
      <div sx={styles.sidebarPageContents}>
        {props.children}
      </div>
    </div>
  )
}
