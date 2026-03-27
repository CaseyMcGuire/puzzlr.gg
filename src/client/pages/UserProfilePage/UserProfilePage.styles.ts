import * as stylex from "@stylexjs/stylex";

export const UserProfileSectionStyles = stylex.create({
  list: {
    display: "flex",
    flexDirection: "column",
    gap: "12px",
  },
  listRow: {
    borderWidth: "1px",
    borderStyle: "solid",
    borderColor: "#e5e7eb",
    borderRadius: "12px",
    padding: "14px",
    backgroundColor: "#f8fafc",
  },
  rowTitle: {
    margin: 0,
    fontSize: "1rem",
    fontWeight: "600",
    wordBreak: "break-word",
  },
  rowMeta: {
    marginTop: "8px",
    display: "flex",
    flexWrap: "wrap",
    gap: "10px",
    color: "#475569",
    fontSize: "0.92rem",
  },
  link: {
    color: "#1d4ed8",
    textDecoration: "none",
    fontWeight: "600",
  },
  emptyState: {
    margin: 0,
    color: "#64748b",
  },
});
