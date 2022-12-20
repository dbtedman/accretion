import { expect, test } from "vitest";
import { mount } from "@vue/test-utils";
import AccretionSearch from "@/components/AccretionSearch/AccretionSearch.vue";

test("displays search icon", () => {
    const wrapper = mount(AccretionSearch, {});
    expect(wrapper.find("[data-test-id='search-icon']").html()).toContain(
        "search"
    );
});

test("displays close icon", () => {
    const wrapper = mount(AccretionSearch, {});
    expect(wrapper.find("[data-test-id='close-icon']").html()).toContain(
        "close"
    );
});
