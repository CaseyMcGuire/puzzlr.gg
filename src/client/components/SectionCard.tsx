import * as stylex from "@stylexjs/stylex";
import {ReactNode} from "react";

type Props = {
  title: string;
  children: ReactNode;
};

const styles = stylex.create({
  card: {
    borderWidth: "1px",
    borderStyle: "solid",
    borderColor: "#d9e1ec",
    borderRadius: "16px",
    backgroundColor: "#ffffff",
    padding: "20px",
    display: "flex",
    flexDirection: "column",
    gap: "16px",
    flex: "1 1 280px",
  },
  title: {
    margin: 0,
    fontSize: "1.2rem",
  },
});

export default function SectionCard({title, children}: Props) {
  return (
    <div sx={styles.card}>
      <h2 sx={styles.title}>{title}</h2>
      {children}
    </div>
  );
}
