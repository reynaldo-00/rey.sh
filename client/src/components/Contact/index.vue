<template>
    <section class="section contact">
        <h2>CONTACT</h2>
        <Socials/>
    </section>
</template>

<script>
import Socials from './Socials';
export default {
    name: "Contact",
    components: {Socials},
}
</script>

<style lang="scss">
    section.contact {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        padding: 48px;
        width: 100vw;
        background: #0e0e0e;
        h2 {
            letter-spacing: 14px;
            font-size: 30px;
            margin: 0;
            margin-bottom: 20px;
        }
    }

</style>
