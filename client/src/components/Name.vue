<template>
    <div>
        <h2>
            {{ first }}
        </h2>
        <h2>
            {{ middle }}
        </h2>
        <h2>
            {{ last }}
        </h2>
    </div>
</template>

<script>
export default {
    name: 'Name',
    data: () => ({
        first: 'CHRISTOPHER',
        middle: 'REYNALDO',
        last: 'CEDENO',
    })
}
</script>

<style scoped>
    div {
        margin-top: 50px;
    }
    h2 {
        font-size: 70px;
        font-weight: 800;
        margin: 0px;
        line-height: 70px;
    }
</style>
