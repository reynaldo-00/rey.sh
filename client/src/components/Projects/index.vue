<template>
    <section class="section projects">
        <h2>PROJECTS</h2>
        <div class="carousel">
            <div class="carousel-item">
                <h3 class="title">FlexLog</h3>
                <img class="preview-img" src="../../assets/flexlog.png"/>
                <p class="description">Fitness tracking app built using React Hooks for state management, Node.js for backend language</p>
            </div>
            <div class="carousel-item">
                <h3 class="title">Google Keep Clone</h3>
                <img class="preview-img" src="../../assets/google-keep.png"/>
                <p class="description">Fitness tracking app built using React Hooks for state management, Node.js for backend language</p>
            </div>
            <div class="carousel-item">
                <h3 class="title">Carbtagropher</h3>
                <img class="preview-img" src="../../assets/carbtographer.png"/>
                <p class="description">Fitness tracking app built using React Hooks for state management, Node.js for backend language</p>
            </div>
            <div class="carousel-item">
                <h3 class="title">rey.sh</h3>
                <img class="preview-img" src="../../assets/rey.png"/>
                <p class="description">Fitness tracking app built using React Hooks for state management, Node.js for backend language</p>
            </div>
        </div>

    </section>
</template>

<script>
export default {
    name: 'Projects',
}
</script>

<style lang="scss">
    section.projects {
        width: 100vw;
        height: auto;
        background: rgb(245,245,245);
        color: #020424;
        padding: 48px;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: flex-start;
        .carousel-item {
            display: flex;
            flex-direction: column;
            justify-content: space-around;
            align-items: center;
            .title {
                font-size: 22px;
            }
            .preview-img {
                height: 320px;
                width:  auto;
            }
            .description {
                font-size: 18px;
            }
        }
    }
</style>
