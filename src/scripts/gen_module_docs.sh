# enable executable: chmod +x ./scripts/<executable_name>.sh
#!/bin/bash
if [ -d "docs" ]; then
  mkdir "docs"
fi
# cp -fr -R . ./docs
# Generate docs: