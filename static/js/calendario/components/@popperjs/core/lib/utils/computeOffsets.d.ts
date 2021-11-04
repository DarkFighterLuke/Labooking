import type {ClientRectObject, Offsets, PositioningStrategy, Rect} from "../types";
import {Placement} from "../enums";

export default function computeOffsets({reference, element, placement}: {
    reference: Rect | ClientRectObject;
    element: Rect | ClientRectObject;
    strategy: PositioningStrategy;
    placement?: Placement;
}): Offsets;
