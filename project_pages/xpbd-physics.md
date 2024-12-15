# XPBD Physics with a Quadtree

This is an implementation of the Extended Position Based Dynamic (XPBD) resolution algorithm as outlined <a href=https://matthias-research.github.io/pages/publications/XPBD.pdf target=_blank>here</a>. The project was developed in Rust, using the Bevy Gameframe work. Each collider is a component attached to an entity and we use systems to query the world for our colliders. From there we can apply velocity, check for collision contacts and resolve those contacts.

To view the Github Repository <a href=https://github.com/colorfulparadox/XPBD-Physics target=_blank>click here</a>

After implementation of XPBD, an issue occurs where the more colliders you add into the world the slower your program will run. The naive solution to check for collision contacts is a O(n^2) algorithm. We can solve this by implementing a data structure that organizes the world into sections to query against.

To solve this problem I used a Quadtree. A Quadtree takes the world space and divides it into four buckets. The buckets represent a square section of the world. As more colliders are added the buckets subdivide into smaller buckets. To query a Quadtree is approximately O(n+m) which changes the time complexity to exponential to a linear time complexity. The sub-division behavior can be observed in the following video.

<iframe width="420" height="315" 
    src="https://www.youtube.com/embed/my8O3WiS_CI" 
    title="Physics Demo Video" 
    frameborder="0" 
    allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" 
    allowfullscreen>
</iframe>
