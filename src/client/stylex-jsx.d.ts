import type {
  CompiledStyles,
  InlineStyles,
  StyleXArray,
} from "@stylexjs/stylex";

type StyleXProp = StyleXArray<
  | (null | undefined | CompiledStyles)
  | boolean
  | Readonly<[CompiledStyles, InlineStyles]>
>;

declare module "react" {
  interface HTMLAttributes<T> {
    sx?: StyleXProp;
  }
  interface SVGAttributes<T> {
    sx?: StyleXProp;
  }
}