+++
title = 'Quirky Things in C++'
date = 2024-07-26T19:45:46-04:00
draft = true
+++

The following post is meant to document the weird things I've encountered when programming in C++. If you're experienced, you'll probably know most of the things on here. However, these are things I've discovered in my C++ journey:


## Vector Initializations: When Random Isn't so Random

Suppose you're making a 2D game where the map consists of tiles. For some reason, say you need each tile to have a random seed. You could design your `Tile` class in the following way:
```cpp
// tile.hpp
class Tile {
public:
    Tile();  // Constructor
private:
    int m_tileSeed;
};
```

```cpp
// tile.cpp
// Constructor
Tile::Tile() {
    m_tileSeed = RandomGenerator::getInt(); // generates a random integer
}
```

Now you want to create a `TileMap` class that holds all `Tile`s in your game:
```cpp
// tile_map.hpp
class TileMap {
public:
    TileMap(std::size_t width, std::size_t height);  // Constructor

private:
    std::vector<std::vector<Tile>> m_tileMap;
    const std::size_t m_width;
    const std::size_t m_height;
};
```

Naturally, here's a constructor you (could) make for your `TileMap`:

```cpp
// tile_map.cpp
TileMap::TileMap(const std::size_t width, const std::size_t height)
    : m_width(width), m_height(height), m_tileMap(height, std::vector<Tile>(width)) {}
```

And now you've got yourself a runtime bug in which all of your `Tile`s have the same seed!

![Debugger view of the TileMap elements](/media/cpp_quirks/repeated_seeds.jpg)

What you need to do instead is initialize the 2D vector _without_ any `Tile`s, and then individually construct each `Tile`:
```cpp
TileMap::TileMap(const std::size_t width, const std::size_t height)
    : m_width(width), m_height(height), m_tileMap(height) {
        for(std::vector<Tile>& row : m_tileMap) {
            row.reserve(width);
            for(std::size_t i = 0; i < width; ++i) {
                row.emplace_back();
            }
        }
}
```

And voila! You've solved the issue:

![Debugger view of the fixed TileMap elements](/media/cpp_quirks/unrepeated_seeds.jpg)

### Understanding `std::vector`'s default constrution behavior
`std::vector` generates a single `Tile`, and then copies that tile across the whole map, which is what leads to every `Tile` having the same seed. If we look inside the [GNU implentation] for `_M_fill_assign`, we can see why this is the case:
```tcc
std::__uninitialized_fill_n_a(this->_M_impl._M_finish, __add, __val, _M_get_Tp_allocator());
```

