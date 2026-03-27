import * as stylex from "@stylexjs/stylex";

const styles = stylex.create({
  hero: {
    display: "flex",
    flexDirection: "column",
    gap: "10px",
    marginBottom: "24px",
    borderWidth: "1px",
    borderStyle: "solid",
    borderColor: "#d9e1ec",
    borderRadius: "20px",
    background: "linear-gradient(135deg, #ffffff 0%, #eef4ff 100%)",
    padding: "28px",
  },
  eyebrow: {
    margin: 0,
    fontSize: "0.82rem",
    fontWeight: "700",
    letterSpacing: "0.08em",
    textTransform: "uppercase",
    color: "#2563eb",
  },
  title: {
    margin: 0,
    fontSize: "2.2rem",
    lineHeight: "1.05",
    wordBreak: "break-word",
  },
  subtitle: {
    margin: 0,
    color: "#4b5563",
    fontSize: "1rem",
  },
});

type Props = {
  title: string;
  subtitle: string;
};

export default function UserProfileHero({title, subtitle}: Props) {
  return (
    <div sx={styles.hero}>
      <p sx={styles.eyebrow}>User profile</p>
      <h1 sx={styles.title}>{title}</h1>
      <p sx={styles.subtitle}>{subtitle}</p>
    </div>
  );
}
