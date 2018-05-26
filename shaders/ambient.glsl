#vertex
#version 330

layout(location = 0) in vec4 position;
layout(location = 1) in vec2 vertTexCoords;
layout(location = 2) in vec3 normals;

out vec2 fragTexCoords;
out vec3 fragPos;
out vec3 fragNormal;

uniform mat4 model;
uniform mat4 projection;
uniform mat4 view;

void main() {
    gl_Position = projection * view * model * position;
    fragPos = vec3(model * position);
    fragTexCoords = vertTexCoords;
    fragNormal = mat3(transpose(inverse(model))) * normals;
}

#fragment
#version 330

in vec3 fragPos;
in vec2 fragTexCoords;
in vec3 fragNormal;

out vec4 result;

struct material {
    sampler2D diffTex;
    sampler2D specTex;
    float shininess;
};

uniform material mat;

void main() { 
    // Color of the texture
    vec4 albedo = texture(mat.diffTex, fragTexCoords);
    // Minimum light.
    vec3 ambient = 0.1 * vec3(texture(mat.diffTex, fragTexCoords));

    result = vec4(ambient, 1.0);
}