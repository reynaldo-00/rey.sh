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
        first: 'Christopher',
        middle: 'Reynaldo',
        last: 'Cedeno',
    })
}
</script>

<style scoped>
    h2 {

    }
</style>
