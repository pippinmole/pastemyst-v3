<script lang="ts">
    import { onMount, type Snippet } from "svelte";
    import { themeStore } from "./stores";
    import { themes, type Theme } from "./themes";

    interface Props {
        currentTheme: string;
        children: Snippet;
    }

    let { currentTheme = $bindable(), children }: Props = $props();

    let mounted = false;

    themeStore.subscribe((theme) => {
        if (!theme) return;

        currentTheme = theme.name;
        if (mounted) setRootColors(theme);
    });

    onMount(() => {
        mounted = true;

        const themeObj = themes.find((t) => t.name === currentTheme) || themes[0];
        themeStore.set(themeObj);
    });

    const setRootColors = (theme: Theme) => {
        for (let [prop, color] of Object.entries(theme.colors)) {
            const varString = `--color-${prop}`;
            document.documentElement.style.setProperty(varString, color);
        }
    };
</script>

<div id="theme-context" data-theme={currentTheme}>
    {@render children()}
</div>
