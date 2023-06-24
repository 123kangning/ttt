pid=$(pgrep -f "qin")
if [[ -n "$pid" ]]; then
    echo "Stopping qin service..."
    kill $pid 
fi
