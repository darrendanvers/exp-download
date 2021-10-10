FROM nginx:latest

#
# Copy the site files.
#
COPY site/* /usr/share/nginx/html

#
# Save the Nginx configuration.
#
COPY default.conf /etc/nginx/conf.d/default.conf