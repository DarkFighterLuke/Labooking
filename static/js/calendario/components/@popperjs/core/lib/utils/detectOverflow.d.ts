import type {Padding, SideObject, State} from "../types";
import type {Boundary, Context, Placement, RootBoundary} from "../enums";

export declare type Options = {
    placement: Placement;
    boundary: Boundary;
    rootBoundary: RootBoundary;
    elementContext: Context;
    altBoundary: boolean;
    padding: Padding;
};
export default function detectOverflow(state: State, options?: Partial<Options>): SideObject;
