import {useState} from "react";
import * as stylex from "@stylexjs/stylex";
import {SidebarStyles} from "components/Sidebar/SidebarStyles.stylex";

type SidebarLinkItem = {
  type: "link";
  label: string;
  href: string;
};

type SidebarFolderItem = {
  type: "folder";
  label: string;
  children: SidebarItem[];
};

type SidebarItem = SidebarLinkItem | SidebarFolderItem;

type SidebarProps = {
  items: SidebarItem[];
};

const styles = stylex.create({
  sidebar: {
    display: 'flex',
    flexDirection: 'column',
    position: 'fixed',
    width: SidebarStyles.sidebarWidth,
    height: '100%',
    borderRightWidth: '1px',
    borderRightStyle: 'solid',
    borderRightColor: 'rgba(92,92,92,0.5)',
  },
  item: {
    padding: '8px 12px',
    cursor: 'pointer',
  },
  link: {
    textDecoration: 'none',
    color: 'inherit',
    display: 'block',
    padding: '8px 12px',
  },
  groupLabel: {
    padding: '8px 12px',
    cursor: 'pointer',
    fontWeight: 'bold',
  },
  children: {
    paddingLeft: '16px',
  },
});

type Props = {
  item: SidebarItem;
}

function SidebarItemView({item}: Props) {
  switch (item.type) {
    case "folder":
      return <SidebarFolder item={item} />;
    case "link":
      return <SidebarLink item={item} />;
    default:
      return null;
  }
}

function SidebarFolder({item}: {item: SidebarFolderItem}) {
  const [expanded, setExpanded] = useState(false);

  return (
    <div>
      <div
        {...stylex.props(styles.groupLabel)}
        onClick={() => setExpanded(!expanded)}
      >
        {item.label}
      </div>
      {expanded && (
        <div {...stylex.props(styles.children)}>
          {item.children.map((child, i) => (
            <SidebarItemView key={i} item={child} />
          ))}
        </div>
      )}
    </div>
  );
}

function SidebarLink({item}: {item: SidebarLinkItem}) {
  return (
    <a href={item.href} {...stylex.props(styles.link)}>
      {item.label}
    </a>
  );
}

export default function Sidebar(props: SidebarProps) {
  return (
    <nav {...stylex.props(styles.sidebar)}>
      {props.items.map((item, i) => (
        <SidebarItemView key={i} item={item} />
      ))}
    </nav>
  );
}
