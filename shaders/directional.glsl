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

struct Material {
    sampler2D diffTex;
    sampler2D specTex;
    float shininess;
};

struct Light {
    float intensity;
    vec3 direction;
};

uniform Light dl;
uniform Material mat;
uniform vec3 viewPos;

void main() { 
    // Diffuse lighting.
    vec3 lightDir = normalize(-dl.direction);
    vec3 norm = normalize(fragNormal);
    float diff = max(dot(norm, lightDir), 0.0);
    vec3 diffuse = dl.intensity * diff * vec3(texture(mat.diffTex, fragTexCoords));

    // Specularity, the shiny effect when right in the light.
    vec3 viewDir = normalize(viewPos - fragPos);
    vec3 reflectDir = reflect(-lightDir, norm);
    float spec = pow(max(dot(viewDir, reflectDir), 0.0), 32);
    vec3 specular = mat.shininess * spec * vec3(texture(mat.specTex, fragTexCoords));

    result = vec4(diffuse + specular, 1.0);
}