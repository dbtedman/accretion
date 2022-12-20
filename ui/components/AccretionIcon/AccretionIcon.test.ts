import { expect, test } from "vitest";
import { mount } from "@vue/test-utils";
import AccretionIcon from "@/components/AccretionIcon/AccretionIcon.vue";

test("can render search icon", () => {
    const wrapper = mount(AccretionIcon, {
        props: {
            kind: "search",
        },
    });
    expect(wrapper.html()).toContain("search");
});

test("can render close icon", () => {
    const wrapper = mount(AccretionIcon, {
        props: {
            kind: "close",
        },
    });
    expect(wrapper.html()).toContain("close");
});
