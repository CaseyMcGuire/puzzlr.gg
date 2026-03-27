import * as stylex from "@stylexjs/stylex";
import {ReactNode} from "react";

type Props = {
  label: string;
  value: ReactNode;
};

const styles = stylex.create({
  card: {
    display: 'flex',
    flexDirection: 'column',
    flexGrow: 1,
    borderWidth: "1px",
    borderStyle: "solid",
    borderColor: "#d9e1ec",
    borderRadius: "16px",
    backgroundColor: "#ffffff",
    padding: "18px",
  },
  label: {
    margin: 0,
    color: "#64748b",
    fontSize: "0.9rem",
  },
  value: {
    marginTop: "8px",
    marginBottom: 0,
    fontSize: "1.8rem",
    fontWeight: "700",
  },
});

export default function KpiCard({label, value}: Props) {
  return (
    <div sx={styles.card}>
      <span sx={styles.label}>{label}</span>
      <span sx={styles.value}>{value}</span>
    </div>
  );
}
