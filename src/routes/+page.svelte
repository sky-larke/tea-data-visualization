<script lang="ts">
    import { echarts } from "$lib/echarts";
    const option = {
        xAxis: {
            type: "category",
            data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
        },
        yAxis: {
            type: "value",
        },
        series: [
            {
                data: [150, 230, 224, 218, 135, 147, 260],
                type: "scatter",
            },
        ],
    };
</script>

<main class="p-4">
    <h1>Tea Data Visualization</h1>
    <h2>Overall</h2>
    <h3>Rank vs. Cost ($/g)</h3>
    <div class="container" use:echarts={option}></div>
    <h3>Rank vs. Price Efficiency (Cost-Rank Score)</h3>
</main>

<style>
    .container {
        width: 500px;
        height: 500px;
    }
</style>
