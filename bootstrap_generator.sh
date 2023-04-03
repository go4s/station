#!/usr/bin/env sh

cat > bootstrap.sh << EOF
#!/usr/bin/env sh

exec /var/app/$MODULE_NAME
EOF

chmod 700 bootstrap.sh
